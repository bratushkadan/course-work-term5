import './App.css';

import { createBrowserRouter, RouterProvider, Link, useParams, Outlet } from 'react-router-dom';
import styled from 'styled-components';
import { LoginPanel } from './components/LoginPanel';
import { NavMenu } from './components/NavMenu';
import { NotFound } from './pages/NotFound';
import { Catalog } from './pages/Catalog';
import { StoresPage } from './pages/Stores';
import { StorePage } from './pages/Store';
import { OrdersPage } from './pages/Orders';
import { FavoritesPage } from './pages/Favorites';
import {OrderPage} from './pages/Order';
import {CartPage} from './pages/Cart';
import {ProductPage} from './pages/Product';
import {RequestAppDataComponent} from './components/RequestAppData';

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
  margin-left: 0.33rem;
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
      <RequestAppDataComponent/>
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
      { path: 'cart', element: <CartPage/> },
      { path: 'stores', element: <StoresPage /> },
      { path: 'stores/:id', element: <StorePage /> },
      { path: 'products/:id', element: <ProductPage /> },
      { path: 'orders', element: <OrdersPage /> },
      { path: 'orders/:id', element: <OrderPage /> },
      { path: 'favorites', element: <FavoritesPage /> },
      { path: '*', element: <NotFound /> },
    ],
  },
]);
