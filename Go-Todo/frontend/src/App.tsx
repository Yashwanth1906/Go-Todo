import { useEffect, useState } from 'react'
import './App.css'
import {GetTasks} from "../wailsjs/go/main/App"
import { models } from "../wailsjs/go/models"

function App() {
  const [todos, setTodos] = useState<models.Task[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    const fetchTasks = async () => {
      try {
        const tasks = await GetTasks();
        console.log('Fetched tasks:', tasks);
        setTodos(tasks);
      } catch (err) {
        console.error('Error fetching tasks:', err);
        setError('Failed to fetch tasks');
      } finally {
        setLoading(false);
      }
    };

    fetchTasks();
  }, []);

  if (loading) return <div>Loading...</div>;
  if (error) return <div>Error: {error}</div>;

  return (
    <>
      <h1>Todo App</h1>
      <div>
        <h2>Tasks ({todos.length})</h2>
        {todos.length === 0 ? (
          <p>No tasks found</p>
        ) : (
          <ul>
            {todos.map((todo) => (
              <li key={todo.id}>
                <strong>{todo.name}</strong> - {todo.description} 
                <span> [{todo.status === 0 ? 'Pending' : 'Completed'}]</span>
              </li>
            ))}
          </ul>
        )}
      </div>
    </>
  )
}
export default App

// const App = () => {
//   return (
//     <>
//       <h1>Hello World</h1>
//     </>
//   )
// }

// export default App