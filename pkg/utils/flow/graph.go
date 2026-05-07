// SPDX-FileCopyrightText: SAP SE or an SAP affiliate company and Gardener contributors
//
// SPDX-License-Identifier: Apache-2.0

package flow

import (
	"fmt"

	"k8s.io/utils/clock"
)

// Task is a unit of work. It has a name, a payload function and a set of dependencies.
// A is only started once all its dependencies have been completed successfully.
type Task struct {
	Name         string
	Fn           TaskFn
	SkipIf       bool
	Dependencies TaskIDs
}

func (t *Task) ID() TaskID {
	return TaskID(t.Name)
}

// Spec returns the TaskSpec of a task.
func (t *Task) Spec() *TaskSpec {
	return &TaskSpec{
		t.Fn,
		t.SkipIf,
		t.Dependencies.Copy(),
	}
}

// TaskSpec is functional body of a Task, consisting only of the payload function and
// the dependencies of the Task.
type TaskSpec struct {
	Fn           TaskFn
	Skip         bool
	Dependencies TaskIDs
}

// Tasks is a mapping from TaskID to TaskSpec.
type Tasks map[TaskID]*TaskSpec

type TaskGroup struct {
	id    TaskID
	tasks map[TaskID]Task

	// dependencies on other tasks or groups.
	dependencies TaskIDs
}

func (g TaskGroup) TaskIDs() []TaskID {
	taskIDs := make(TaskIDs, len(g.tasks))
	for id := range g.tasks {
		taskIDs.Insert(id)
	}
	return taskIDs.TaskIDs()
}

func NewTaskGroup(id TaskID, tasks ...Task) TaskGroup {
	return (TaskGroup{
		id:           id,
		tasks:        make(map[TaskID]Task, len(tasks)),
		dependencies: make(TaskIDs, len(tasks)),
	}).AddAll(tasks...)
}

func (g TaskGroup) AddAll(tasks ...Task) TaskGroup {
	for _, task := range tasks {
		g.Add(task)
	}
	return g
}

func (g TaskGroup) Add(task Task) TaskID {
	id := task.ID()
	if _, ok := g.tasks[id]; ok {
		panic(fmt.Sprintf("Task with id %q already exists in group %q", id, g.id))
	}
	g.tasks[id] = task
	return id
}

func (g TaskGroup) WithDependencies(dependencies ...TaskIDer) TaskGroup {
	for _, dependency := range dependencies {
		g.dependencies.Insert(dependency)
	}
	return g
}

// Graph is a builder for a Flow.
type Graph struct {
	name   string
	tasks  Tasks
	groups map[TaskID]TaskGroup

	// Clock is used to retrieve the current time.
	Clock clock.Clock
}

// Name returns the name of a graph.
func (g *Graph) Name() string {
	return g.name
}

// NewGraph returns a new Graph with the given name.
func NewGraph(name string) *Graph {
	return &Graph{name: name, tasks: make(Tasks), groups: make(map[TaskID]TaskGroup), Clock: clock.RealClock{}}
}

// Add adds the given Task to the graph.
// This panics if
// - There is already a Task present with the same name
// - One of the dependencies of the Task is not present
func (g *Graph) Add(task Task) TaskID {
	id := task.ID()
	if _, ok := g.tasks[id]; ok {
		panic(fmt.Sprintf("Task with id %q already exists", id))
	}

	spec := task.Spec()
	for dependencyID := range spec.Dependencies {
		if _, ok := g.tasks[dependencyID]; !ok {
			panic(fmt.Sprintf("Task %q is missing dependency %q", id, dependencyID))
		}
	}
	g.tasks[id] = task.Spec()
	return id
}

func (g *Graph) AddGroup(group TaskGroup) TaskIDs {
	g.groups[group.id] = group

	dependencies := make(TaskIDs)
	for dependency := range group.dependencies {
		if gg, isGroup := g.groups[dependency]; isGroup {
			dependencies.Insert(gg)
		} else {
			dependencies.Insert(dependency)
		}
	}

	ids := make(TaskIDs, len(g.tasks))
	for _, task := range group.tasks {
		if task.Dependencies == nil {
			task.Dependencies = make(TaskIDs, len(dependencies))
		}
		task.Dependencies.Insert(dependencies)
		ids.Insert(g.Add(task))
	}

	return ids
}

// Compile compiles the graph into an executable Flow.
func (g *Graph) Compile() *Flow {
	nodes := make(nodes, len(g.tasks))

	for taskName, taskSpec := range g.tasks {
		for dependencyID := range taskSpec.Dependencies {
			dependency := nodes.getOrCreate(dependencyID)
			dependency.addTargets(taskName)
		}

		node := nodes.getOrCreate(taskName)
		node.fn = taskSpec.Fn
		node.skip = taskSpec.Skip
		node.required = taskSpec.Dependencies.Len()
	}

	return &Flow{
		name:  g.name,
		nodes: nodes,
		clock: g.Clock,
	}
}
