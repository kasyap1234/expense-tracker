// components/AddExpense.tsx

export interface Expense {
    id : number; 
    title : string ; 
    amount : number ; 

}
export default function AddExpense() {
    const handleSubmit = async (e) => {
      e.preventDefault()
      const token = localStorage.getItem('token')
      // Submit expense to API
    }
    
    return (
      <form onSubmit={handleSubmit}>
        <input name="title" placeholder="Expense Title" />
        <input name="amount" type="number" placeholder="Amount" />
        <button type="submit">Add Expense</button>
      </form>
    )
  }
  

  