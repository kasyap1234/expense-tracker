"use client"
import { useEffect, useState } from 'react'
import { Expense } from '@/component/AddExpense'
import AddExpense from '@/component/AddExpense'

export default function Dashboard() {
  const [expenses, setExpenses] = useState<Expense[]>([])

  const fetchExpenses = () => {
    const token = localStorage.getItem('token')
    fetch('http://localhost:8000/expenses', {
      headers: {
        'Authorization': `Bearer ${token}`
      }
    })
    .then(res => res.json())
    .then(data => setExpenses(data))
  }

  useEffect(() => {
    fetchExpenses()
  }, [])

  return (
    <div className="p-8">
      <h1 className="text-2xl font-bold mb-4">Your Expenses</h1>
      <AddExpense />
      <div className="grid gap-4 mt-8">
        {expenses.map((expense: Expense) => (
          <div key={expense.id} className="border p-4 rounded">
            <h2>{expense.title}</h2>
            <p>${expense.amount}</p>
          </div>
        ))}
      </div>
    </div>
  )
}


