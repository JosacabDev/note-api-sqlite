package tasks

import (
	"fmt"
	customErrors "github/JosacabDev/api-sqlite/pkg/errors"
	"github/JosacabDev/api-sqlite/pkg/libjson"
	"log"
	"net/http"
)

type HandlerTask struct {
	Repo Repository
}

func NewHandlerTask(repo Repository) *HandlerTask {
	return &HandlerTask{
		Repo: repo,
	}
}

func (h *HandlerTask) GetAllTasks(w http.ResponseWriter, r *http.Request) {
	tasks, err := h.Repo.GetAllTasks()
	if err != nil {
		log.Println("Failed to retrieve tasks:", err)
		libjson.EncodeCustomError(w, customErrors.InternalError(fmt.Sprintf("Failed to retrieve tasks: %s.", err.Error())))
		return
	}

	libjson.EncodeOk(w, tasks)
}
