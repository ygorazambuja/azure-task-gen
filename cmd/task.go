package cmd

type Task struct {
	ID            string
	WorkItemType  string
	Title         string
	AssignedTo    string
	State         string
	AreaID        string
	IterationID   string
	ItemContrato  string
	IDSPF         string
	UST           string
	Complexidade  string
	Activity      string
	Description   string
	EstimateMade  string
	RemainingWork string
}

type TaskList struct {
	Tasks []Task
}

func (t TaskList) GetTasks() []Task {
	return t.Tasks
}

func (t TaskList) GenerateCSV() string {
	csv := "ID,Work Item Type,Title,Assigned To,State,Area ID,Iteration ID,Item Contrato,ID SPF,UST,Complexidade,Activity,Description,Estimate Made,Remaining Work\n"

	for _, task := range t.Tasks {
		csv += task.ID + "," +
			task.WorkItemType + "," +
			task.Title + "," +
			task.AssignedTo + "," +
			task.State + "," +
			task.AreaID + "," +
			task.IterationID + "," +
			task.ItemContrato + "," +
			task.IDSPF + "," +
			task.UST + "," +
			task.Complexidade + "," +
			task.Activity + "," +
			task.Description + "," +
			task.EstimateMade + "," +
			task.RemainingWork + "\n"
	}

	return csv
}
