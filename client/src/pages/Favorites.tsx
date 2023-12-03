import { useFavorite } from '../stores';
import { useShallow } from 'zustand/react/shallow';
import { Link } from 'react-router-dom';
import { FavoriteProductComponent } from '../components/FavoriteProduct';

export const FavoritesPage: React.FC = () => {
  const { favorites, setFavorites } = useFavorite(
    useShallow((state) => ({
      favorites: state.favorites,
      setFavorites: state.setFavorites,
    }))
  );

  return (
    <>
      <h2>
        {favorites.length === 0 ? (
          <>
            Нет товаров, добавленных в избранное. <Link to="/">Добавим?</Link>
          </>
        ) : (
          'Избранное'
        )}
      </h2>
      {favorites.map((props) => (
        <FavoriteProductComponent {...props} key={props.product_id} isWithProductLink={true} />
      ))}
    </>
  );
};
