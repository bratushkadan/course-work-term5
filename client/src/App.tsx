import './App.css'

import { createBrowserRouter, RouterProvider, Link, useParams } from 'react-router-dom';
import styled from 'styled-components';
import {LoginPanel} from './components/LoginPanel';

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
      element: null,
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

const PageWrapper = styled.div`
  display: flex;
  flex-direction: column;
  height: 100vh;
`

const HeaderContent = styled.header`
  min-height: 1rem;
  background-color: #d3d3d3;
  display: flex;
  flex-direction: row;
  justify-content: space-between;
`

const HeaderLogo = styled.h2`
  font-size: 2.25rem;
  color: #0d2901;
  margin: 0 0 0 1rem;
`

const PageContent = styled.main`
  flex: 1;
  flex-grow: 1;
  overflow-y: auto;
  
`

const FooterContent = styled.footer`
  min-height: 3rem;
  background-color: #d3d3d3;
  display: flex;
  position: sticky;
  justify-content: center;
  align-items: center;
`

function App() {
  return (
    <PageWrapper>
      <HeaderContent>
        <HeaderLogo>Floral</HeaderLogo>
        <LoginPanel></LoginPanel>
      </HeaderContent>
      <PageContent>
        <RouterProvider router={router} />
      </PageContent>
      <FooterContent>
        <b>Floral</b><span> — Данила Братушка 2023</span> &copy;
      </FooterContent>
    </PageWrapper>
  )
}

export default App
