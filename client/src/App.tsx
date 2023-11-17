import { useState } from 'react'
import reactLogo from './assets/react.svg'
import viteLogo from '/vite.svg'
import './App.css'

import { createBrowserRouter, RouterProvider, Link, useParams } from 'react-router-dom';

function Page() {
  const {id} = useParams()

  return <>
    <Link to="/"><h1>Go to main page</h1></Link>
    <h1>Todo {id}</h1>
  </>;
}

// https://reactrouter.com/en/main/routers/create-browser-router#routes
const router = createBrowserRouter(
  [
    {
      path: '/',
      element: <DefaultPage />,
    },
    {
      path: '/todo/:id',
      element: <Page/>,
    },
    {
      path: '*',
      element: (
        <>
          <h1>Not Found :(</h1>
        </>
      ),
    },
  ]
);

function DefaultPage() {
  const [count, setCount] = useState(0);

  return (
    <>
    <div>
        <a href="https://vitejs.dev" target="_blank">
          <img src={viteLogo} className="logo" alt="Vite logo" />
        </a>
        <a href="https://react.dev" target="_blank">
          <img src={reactLogo} className="logo react" alt="React logo" />
        </a>
    </div>
      <h1>Vite + React</h1>
      <div className="card">
        <button onClick={() => setCount((count) => count + 1)}>
          count is {count}
        </button>
        <p>
          Edit <code>src/App.tsx</code> and save to test HMR
        </p>
      </div>
      <p className="read-the-docs">
        Click on the Vite and React logos to learn more
      </p>
    </>
  )
}

function App() {
  return (
    <RouterProvider router={router} />
  )
}

export default App
