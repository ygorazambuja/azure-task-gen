package cmd

import (
	"fmt"
	"strconv"

	"github.com/ygorazambuja/azure-task-gen/config"
)

type Task struct {
	ID               string
	WorkItemType     string
	Title            string
	AssignedTo       string
	State            string
	AreaID           int
	IterationID      int
	ItemContrato     string
	IDSPF            int
	UST              int
	Complexidade     string
	Activity         string
	Description      string
	EstimateMade     int
	RemainingWork    int
	OriginalEstimate int
}

type TaskList struct {
	Tasks []Task
}

func NewTask(
	id string,
	workItemType string,
	title string,
	assignedTo string,
	state string,
	areaID int,
	iterationID int,
	itemContrato string,
	idSPF int,
	ust int,
	complexidade string,
	activity string,
	description string,
	estimateMade int,
	remainingWork int,
) Task {
	return Task{
		ID:            id,
		WorkItemType:  workItemType,
		Title:         title,
		AssignedTo:    assignedTo,
		State:         state,
		AreaID:        areaID,
		IterationID:   iterationID,
		ItemContrato:  itemContrato,
		IDSPF:         idSPF,
		UST:           ust,
		Complexidade:  complexidade,
		Activity:      activity,
		Description:   description,
		EstimateMade:  estimateMade,
		RemainingWork: remainingWork,
	}
}

func NewTaskList(tasks []Task) TaskList {
	return TaskList{
		Tasks: tasks,
	}
}

func (t TaskList) GetTasks() []Task {
	return t.Tasks
}

func (t TaskList) GenerateCSV() string {
	csv := "ID,Work Item Type,Title,Assigned To,State,Area ID,Iteration ID,Item Contrato,ID SPF,UST,Complexidade,Activity,Description,Estimate Made,Remaining Work,Original Estimate\n"
	for _, task := range t.Tasks {
		csv += task.ID + "," +
			task.WorkItemType + "," +
			task.Title + "," +
			buildAssignedToTask() + "," +
			task.State + "," +
			strconv.Itoa(task.AreaID) + "," +
			strconv.Itoa(task.IterationID) + "," +
			task.ItemContrato + "," +
			strconv.Itoa(task.IDSPF) + "," +
			strconv.Itoa(task.UST) + "," +
			task.Complexidade + "," +
			task.Activity + "," +
			task.Description + "," +
			strconv.Itoa(task.EstimateMade) + "," +
			strconv.Itoa(task.RemainingWork) + "," +
			strconv.Itoa(task.OriginalEstimate) + "\n"
	}

	return csv
}

func buildAssignedToTask() string {
	return fmt.Sprintf("%s <%s>", config.AppConfig.DefaultTask.AssignedTo, config.AppConfig.DefaultTask.Email)
}
