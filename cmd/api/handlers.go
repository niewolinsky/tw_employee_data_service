package main

import (
	"math"

	data "github.com/niewolinsky/tw_employee_data_service/data"
	util "github.com/niewolinsky/tw_employee_data_service/utils"

	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func (app *application) hdlGetEmployee(w http.ResponseWriter, r *http.Request) {
	place_id_str := chi.URLParam(r, "place_id")
	place_id, err := strconv.Atoi(place_id_str)
	if err != nil {
		util.BadRequestResponse(w, r, err)
		return
	}

	employee, err := app.data_access.GetEmployee(r.Context(), int32(place_id))
	if err != nil {
		util.ServerErrorResponse(w, r, err)
		return
	}

	err = util.WriteJSON(w, http.StatusOK, util.Wrap{"employee": employee}, nil)
	if err != nil {
		util.ServerErrorResponse(w, r, err)
		return
	}
}

func (app *application) hdlListEmployee(w http.ResponseWriter, r *http.Request) {
	limit_str := r.URL.Query().Get("limit")
	offset_str := r.URL.Query().Get("offset")

	var limit int32 = math.MaxInt32
	var offset int32 = 0

	if limit_str != "" {
		limitParsed, err := strconv.ParseInt(limit_str, 10, 32)
		if err == nil {
			limit = int32(limitParsed)
		}
	}
	if offset_str != "" {
		offsetParsed, err := strconv.ParseInt(offset_str, 10, 32)
		if err == nil {
			offset = int32(offsetParsed)
		}
	}

	input := data.ListEmployeesParams{
		Limit:  limit,
		Offset: offset,
	}

	employees, err := app.data_access.ListEmployees(r.Context(), input)
	if err != nil {
		util.ServerErrorResponse(w, r, err)
		return
	}

	err = util.WriteJSON(w, http.StatusOK, util.Wrap{"employees": employees}, nil)
	if err != nil {
		util.ServerErrorResponse(w, r, err)
		return
	}
}

func (app *application) hdlPostEmployee(w http.ResponseWriter, r *http.Request) {
	input := data.CreateEmployeeParams{}

	err := util.ReadJSON(w, r, &input)
	if err != nil {
		util.BadRequestResponse(w, r, err)
		return
	}

	err = app.validator.Struct(input)
	if err != nil {
		util.FailedValidationResponse(w, r, err)
		return
	}

	result, err := app.data_access.CreateEmployee(r.Context(), input)
	if err != nil {
		util.ServerErrorResponse(w, r, err)
		return
	}

	result_id, err := result.LastInsertId()
	if err != nil {
		util.ServerErrorResponse(w, r, err)
		return
	}

	employee, err := app.data_access.GetEmployee(r.Context(), int32(result_id))
	if err != nil {
		util.ServerErrorResponse(w, r, err)
		return
	}

	err = util.WriteJSON(w, http.StatusOK, util.Wrap{"employee": employee}, nil)
	if err != nil {
		util.ServerErrorResponse(w, r, err)
		return
	}
}

func (app *application) hdlPutEmployee(w http.ResponseWriter, r *http.Request) {
	employee_id_str := chi.URLParam(r, "employee_id")
	employee_id, err := strconv.Atoi(employee_id_str)
	if err != nil {
		util.BadRequestResponse(w, r, err)
		return
	}

	_, err = app.data_access.GetEmployee(r.Context(), int32(employee_id))
	if err != nil {
		util.ServerErrorResponse(w, r, err)
		return
	}

	input := data.UpdateEmployeeParams{}
	err = util.ReadJSON(w, r, &input)
	if err != nil {
		util.BadRequestResponse(w, r, err)
		return
	}

	input.EmployeeID = int32(employee_id)

	err = app.validator.Struct(input)
	if err != nil {
		util.FailedValidationResponse(w, r, err)
		return
	}

	result, err := app.data_access.UpdateEmployee(r.Context(), input)
	if err != nil {
		util.ServerErrorResponse(w, r, err)
		return
	}

	result_id, err := result.LastInsertId()
	if err != nil {
		util.ServerErrorResponse(w, r, err)
		return
	}

	employee, err := app.data_access.GetEmployee(r.Context(), int32(result_id))
	if err != nil {
		util.ServerErrorResponse(w, r, err)
		return
	}

	err = util.WriteJSON(w, http.StatusOK, util.Wrap{"employee": employee}, nil)
	if err != nil {
		util.ServerErrorResponse(w, r, err)
		return
	}
}

func (app *application) hdlDeleteEmployee(w http.ResponseWriter, r *http.Request) {
	employee_id_str := chi.URLParam(r, "employee_id")
	employee_id, err := strconv.Atoi(employee_id_str)
	if err != nil {
		util.BadRequestResponse(w, r, err)
		return
	}

	err = app.data_access.DeleteEmployee(r.Context(), int32(employee_id))
	if err != nil {
		util.ServerErrorResponse(w, r, err)
		return
	}

	err = util.WriteJSON(w, http.StatusOK, util.Wrap{"deleted": "ok"}, nil)
	if err != nil {
		util.ServerErrorResponse(w, r, err)
		return
	}
}
