"use client"
import { useEffect, useState } from 'react'

interface Expense {
  ID: number
  Title: string
  Description: string
  Amount: number
  UserID: number
}

export default function Dashboard() {
  const [expenses, setExpenses] = useState<Expense[]>([])
  const [title, setTitle] = useState('')
  const [description, setDescription] = useState('')
  const [amount, setAmount] = useState('')

  useEffect(() => {
    fetchExpenses()
  }, [])

  const fetchExpenses = () => {
    fetch('http://localhost:8000/expenses', {
      credentials: 'include',
      headers: {
        'Content-Type': 'application/json'
      }
    })
    .then(res => res.json())
    .then(data => setExpenses(data))
    .catch(err => console.error('Error fetching expenses:', err))
  }
  const handleDelete = async (expenseId: number) => {
    try {
      const response = await fetch(`http://localhost:8000/expenses/${expenseId}`, {
        method: 'DELETE',
        credentials: 'include',
        headers: {
          'Content-Type': 'application/json'
        }
      });
  
      if (response.ok) {
        // Optimistically update UI
       setExpenses(currentExpenses => currentExpenses.filter(expense => expense.ID !==expenseId))
        // Then fetch fresh data
        console.log(expenseId)
        fetchExpenses();
      } else {
        throw new Error('Failed to delete expense');
      }
    } catch (error) {
      console.error('Error deleting expense:', error);
    }
  };
  
  
  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault()
    
    try {
      const response = await fetch('http://localhost:8000/expenses', {
        method: 'POST',
        credentials: 'include',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify({
          Title: title,
          Description: description,
          Amount: parseFloat(amount)
        })
      })
      
      if (response.ok) {
        setTitle('')
        setDescription('')
        setAmount('')
        fetchExpenses()
      }
    } catch (error) {
      console.error('Error adding expense:', error)
    }
  }

  return (
    <div className="max-w-4xl mx-auto p-8">
      <h1 className="text-3xl font-bold mb-8">Expense Dashboard</h1>
      
      {/* Add Expense Form */}
      <div className="bg-white p-6 rounded-lg shadow-md mb-8">
        <h2 className="text-xl font-semibold mb-4">Add New Expense</h2>
        <form onSubmit={handleSubmit} className="space-y-4">
          <div>
            <label className="block text-sm font-medium mb-2">Title</label>
            <input
              type="text"
              value={title}
              onChange={(e) => setTitle(e.target.value)}
              className="w-full p-2 border rounded-md"
              required
            />
          </div>
          <div>
            <label className="block text-sm font-medium mb-2">Description</label>
            <input
              type="text"
              value={description}
              onChange={(e) => setDescription(e.target.value)}
              className="w-full p-2 border rounded-md"
              required
            />
          </div>
          <div>
            <label className="block text-sm font-medium mb-2">Amount</label>
            <input
              type="number"
              value={amount}
              onChange={(e) => setAmount(e.target.value)}
              className="w-full p-2 border rounded-md"
              required
            />
          </div>
          <button 
            type="submit"
            className="w-full bg-blue-500 text-white py-2 px-4 rounded-md hover:bg-blue-600 transition-colors"
          >
            Add Expense
          </button>
        </form>
      </div>

      {/* Expenses List */}
      <div className="bg-white p-6 rounded-lg shadow-md">
        <h2 className="text-xl font-semibold mb-4">Your Expenses</h2>
        <div className="space-y-4">
          {expenses.length === 0 ? (
            <p className="text-gray-500">No expenses yet</p>
          ) : (
            expenses.map((expense) => (
              <div 
                key={expense.ID} 
                className="flex justify-between items-center p-4 border rounded-md hover:bg-gray-50"
              >
                <div>
                  <h3 className="font-medium">{expense.Title}</h3>
                  <p className="text-sm text-gray-500">{expense.Description}</p>
                </div>
                <p className="font-semibold text-lg">
                  ${expense.Amount.toFixed(2)}
                </p>
                <button 
  onClick={() => handleDelete(expense.ID)} 
  className="bg-red-500 text-white px-4 py-2 rounded-md hover:bg-red-600 transition-colors"
>
  Delete
</button>

              </div>
            ))
          )}
        </div>
      </div>
    </div>
  )
}
