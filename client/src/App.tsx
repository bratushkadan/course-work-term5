import './App.css';

import { createBrowserRouter, RouterProvider, Link, useParams, Outlet } from 'react-router-dom';
import styled from 'styled-components';
import { LoginPanel } from './components/LoginPanel';
import { NavMenu } from './components/NavMenu';
import { NotFound } from './pages/NotFound';
import { Catalog } from './pages/Catalog';
import {Stores} from './pages/Stores';
import {Store} from './pages/Store';

const PageWrapper = styled.div`
  display: flex;
  flex-direction: column;
  height: 100vh;
`;

const HeaderContent = styled.header`
  min-height: 1rem;
  background-color: #d3d3d3;
  display: flex;
  flex-direction: row;
  justify-content: space-between;
`;

const HeaderLogo = styled.h2`
  font-size: 2.25rem;
  color: #0d2901;
  margin: 0 0 0 1rem;
  flex-grow: 1;
`;

const PageContent = styled.main`
  flex: 1;
  flex-grow: 1;
  overflow-y: auto;
  margin-left: .33rem;
`;

const FooterContent = styled.footer`
  min-height: 3rem;
  background-color: #d3d3d3;
  display: flex;
  position: sticky;
  justify-content: center;
  align-items: center;
`;

const Root: React.FC<React.PropsWithChildren> = () => {
  return (
    <PageWrapper>
      <HeaderContent>
        <HeaderLogo>Floral</HeaderLogo>
        <NavMenu />
        <LoginPanel></LoginPanel>
      </HeaderContent>
      <PageContent>
        <Outlet />
      </PageContent>
      <FooterContent>
        <b>Floral</b>
        <span> — Данила Братушка 2023</span> &copy;
      </FooterContent>
    </PageWrapper>
  );
};

// https://reactrouter.com/en/main/routers/create-browser-router#routes
export const router = createBrowserRouter([
  {
    path: '/',
    element: <Root />,
    children: [
      { path: '', element: <Catalog /> },
      { path: 'cart', element: null },
      { path: 'stores', element: <Stores/>},
      { path: 'stores/:id', element: <Store/>},
      { path: '*', element: <NotFound /> },
    ],
  },
]);
