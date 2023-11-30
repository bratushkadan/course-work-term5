import type React from 'react';
import { NavLink } from 'react-router-dom';
import styled from 'styled-components';
import { useAuth } from '../stores';

const NavMenuWrapper = styled.nav`
  display: flex;
  justify-content: space-around;
  align-items: center;
  flex-grow: 5;

  & a {
    font-size: 1.33rem;
    margin: 0 0.75rem;
  
    cursor: pointer;
  
    transition: 0.125s color;

    &:hover {
      color: #835dff;
    }
  }

  & a.active {
    color: #835dff;
    font-weight: bold;
    &:hover {
      color: inherit;
    }
  }
`;

const MenuItem = styled.div<React.PropsWithChildren<{ isActive?: boolean }>>`
`;

export const NavMenu: React.FC = () => {
  const token = useAuth((state) => state.token);

  return (
    <NavMenuWrapper>
      <NavLink to="/" className={({isActive}) => isActive ? "active" : ""}>Каталог</NavLink>
      <NavLink to="/stores" className={({isActive}) => isActive ? "active" : ""}>Магазины</NavLink>
      
      {token && (
        <>
          <NavLink to="/favorites" className={({isActive}) => isActive ? "active" : ""}>Избранное</NavLink>
          <NavLink to="/cart" className={({isActive}) => isActive ? "active" : ""}>Корзина</NavLink>
          <NavLink to="/orders" className={({isActive}) => isActive ? "active" : ""}>Заказы</NavLink>
        </>
      )}
      {/* <MenuItem>09</MenuItem>
    <MenuItem>10</MenuItem>
    <MenuItem>11</MenuItem>
    <MenuItem>12</MenuItem>
    <MenuItem>13</MenuItem>
    <MenuItem>14</MenuItem>
    <MenuItem>15</MenuItem>
    <MenuItem>16</MenuItem>
    <MenuItem>17</MenuItem>
    <MenuItem>18</MenuItem>
    <MenuItem>19</MenuItem>
    <MenuItem>20</MenuItem> */}
    </NavMenuWrapper>
  );
};
