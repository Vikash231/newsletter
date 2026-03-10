import { useState, useEffect, useCallback } from 'react'
import axios from 'axios'

const API_BASE = '/api'

export function useNews(category) {
  const [articles, setArticles] = useState([])
  const [loading, setLoading] = useState(true)
  const [error, setError] = useState(null)
  const [page, setPage] = useState(1)
  const [hasMore, setHasMore] = useState(true)
  const LIMIT = 10

  const fetchNews = useCallback(async (pg = 1, cat = category) => {
    try {
      setLoading(true)
      setError(null)
      const params = { page: pg, limit: LIMIT }
      if (cat && cat !== 'all') params.category = cat
      const res = await axios.get(`${API_BASE}/news`, { params })
      const { data, total } = res.data
      if (pg === 1) {
        setArticles(data || [])
      } else {
        setArticles(prev => [...prev, ...(data || [])])
      }
      setHasMore((pg * LIMIT) < total)
      setPage(pg)
    } catch (err) {
      setError(err.response?.data?.error || 'Failed to fetch news')
    } finally {
      setLoading(false)
    }
  }, [category])

  useEffect(() => {
    setPage(1)
    setArticles([])
    fetchNews(1, category)
  }, [category])

  const loadMore = useCallback(() => {
    if (!loading && hasMore) {
      fetchNews(page + 1, category)
    }
  }, [loading, hasMore, page, category, fetchNews])

  const refresh = useCallback(() => {
    setPage(1)
    setArticles([])
    fetchNews(1, category)
  }, [category, fetchNews])

  return { articles, loading, error, hasMore, loadMore, refresh }
}

export function useCategories() {
  const [categories, setCategories] = useState(['all'])
  const [loading, setLoading] = useState(true)

  useEffect(() => {
    axios.get(`${API_BASE}/categories`)
      .then(res => {
        setCategories(['all', ...(res.data.data || [])])
      })
      .catch(() => {})
      .finally(() => setLoading(false))
  }, [])

  return { categories, loading }
}
