import { useEffect } from 'react'
import './App.css'
import {GetTasks} from "../wailsjs/go/main/App"

// type Task struct {
// 	ID          int    `json:"id"`
// 	Name        string `json:"name"`
// 	Description string `json:"description"`
// 	Status      Status `json:"status"`
// }

// enum Status {
//   "Completed",
//   "Pending"
// }

// type Task = {
//   ID : number,
//   Name : string,
//   Description : string,
//   Status : Status
// }

function App() {
  useEffect(()=> {
    var x = GetTasks()
    console.log(x)
  },[])
  return (
    <>
        <p>Hello world</p>
    </>
  )
}

export default App
