import React from 'react';
import logo from './logo.svg';
import './App.scss';

import { apiPathBasename } from './api';
import { DOMAIN_PATH_BASENAME } from './constants';

import { createBrowserRouter, RouterProvider, Link, useParams } from 'react-router-dom';

function Page() {
  const {id} = useParams()
  const [todoText, setTodoText] = React.useState("")
  const [isError, setIsError] = React.useState(false)

  React.useEffect(() => {
    fetch(`${apiPathBasename}/todo/${id}`)
      .then(response => response.json())
      .then(data => {
        setTodoText(data.text)
        setIsError(false)
      })
      .catch(error => {
        console.error(error)
        setIsError(true)
      })
  }, [id])

  return <>
    <Link to="/"><h1>Go to main page</h1></Link>
    <h1>Todo {id}</h1>
    {isError ? <h2 style={{color: 'red'}}>Unknown request error occurred.</h2> : <></>}
    <p>{todoText}</p>
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
  ],
  {
    basename: DOMAIN_PATH_BASENAME,
  }
);

function DefaultPage() {
  return (
    <div className="App">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <p>
          Edit <code>src/App.tsx</code> and save to reload.
        </p>
        <a className="App-link" href="https://reactjs.org" target="_blank" rel="noopener noreferrer">
          Learn React
        </a>
      </header>
    </div>
  );
}

function App() {
  return <RouterProvider router={router} />;
}

export default App;
