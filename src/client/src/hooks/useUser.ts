import { useQuery } from 'react-query';
import { User } from 'types';

const fetcher = () =>
  fetch('/api/auth', {
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

export function useUser() {
  const { isLoading, error, data } = useQuery('user', fetcher);
  // if data is not defined, the query has not completed
  let user: User | null = null;
  if (!data?.hasOwnProperty('error')) user = data;
  return { user, isLoading, error };
}
