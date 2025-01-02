  // components/ExpenseList.tsx
interface Expense {
    id : number; 
    title : string ; 
    amount : number ; 

}

function ExpenseList({ expenses }: { expenses: Expense[] }) {
    return <div>{expenses.map(expense => expense.title)}</div>
  }