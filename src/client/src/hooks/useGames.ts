import { useInfiniteQuery, useQuery } from "react-query";

interface urlParam {
  name: string;
  value: string;
}

function buildUrl(base: string, params: urlParam[]): string {
  let url = base;

  for (let i = 0; i < params.length; i++) {
    let p = params[i];
    let symbol = i === 0 ? "?" : "&";
    url += `${symbol}${p.name}=${p.value}`;
  }

  return url;
}

const fetcher = ({ pageParam = "", nameParam = "" }) => {
  const base = "/api/games";

  let url = buildUrl(base, [
    { name: "cursor", value: pageParam },
    { name: "name", value: nameParam },
  ]);

  return fetch(url, { method: "GET", credentials: "include" }).then((res) => {
    if (res.ok) return res.json();
    return { error: true };
  });
};

export function useGames({ nameParam = "" }) {
  const { isLoading, error, data, hasNextPage, fetchNextPage } =
    useInfiniteQuery(["games", nameParam], () => fetcher({ nameParam }), {
      getNextPageParam: (lastPage, pages) => lastPage.NextCursor,
    });
  // if data is not defined, the query has not completed
  return { data, isLoading, error, hasNextPage, fetchNextPage };
}
