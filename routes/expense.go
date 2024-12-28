package routes

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/kasyap1234/expense-tracker/config"
	"github.com/kasyap1234/expense-tracker/models"
)



func CreateExpense(w http.ResponseWriter, r *http.Request){
	userID := r.Context().Value("userID").(uint)

	var expense models.Expense
	json.NewDecoder(r.Body).Decode(&expense)
	expense.UserID=userID ; 
	config.DB.Create(&expense)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(expense)

}

func GetExpenses(w http.ResponseWriter, r* http.Request){
	userID :=r.Context().Value("userID").(uint)
	var expenses []models.Expense 
	config.DB.Where("user_id = ?",userID).Find(&expenses)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(expenses)

}

func GetExpense(w http.ResponseWriter, r *http.Request){
	userID :=r.Context().Value("userID").(uint)
	var expense models.Expense
	expenseID :=expense.ID
	

	config.DB.Where("id= ?  AND user_id = ?",expenseID,userID).First(&expense)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(expense)

}

func UpdateExpense(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(uint)
	expenseID, _ := strconv.Atoi(r.URL.Query().Get("ID"))
	var expense models.Expense
	config.DB.Where("id = ? AND user_id = ?", expenseID, userID).First(&expense)
	json.NewDecoder(r.Body).Decode(&expense)
	config.DB.Save(&expense)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(expense)
}

