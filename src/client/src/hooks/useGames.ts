import { useQuery } from 'react-query';

const fetcher = () =>
  fetch('/api/games', {
    method: 'GET',
    credentials: 'include',
  }).then((res) => {
    if (!res.ok) {
      return {
        error: true
      }
    }
    return res.json()
  });

export function useGames() {
  const { isLoading, error, data } = useQuery('games', fetcher);
  // if data is not defined, the query has not completed
  let games: any[] = data ? data : []
  return { games, isLoading, error };
}
