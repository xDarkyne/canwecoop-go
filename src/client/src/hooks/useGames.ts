import { useInfiniteQuery, useQuery } from 'react-query';

/**
 * 
 * @param cursor 
 * 
 * fetch('/api/games?cursor='+cursor, {
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
 */

const fetcher = ({pageParam = ""}) => {
  const base = '/api/games';
  const url = pageParam ? `${base}?cursor=${pageParam}` : base;
  return fetch(url, { method: 'GET', credentials: 'include' })
    .then((res) => {
      if (res.ok) return res.json();
      return { error: true }
    })
}
  

export function useGames() {
  const { isLoading, error, data, hasNextPage, fetchNextPage } = useInfiniteQuery('games', fetcher, { getNextPageParam: (lastPage, pages) => lastPage.NextCursor, });
  // if data is not defined, the query has not completed
  return { data, isLoading, error, hasNextPage, fetchNextPage };
}
