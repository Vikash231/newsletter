import React, { useEffect, useRef } from 'react'
import NewsCard from './NewsCard.jsx'
import './NewsFeed.css'

function Skeleton() {
  return (
    <div className="skeleton-card">
      <div className="skeleton-img shimmer" />
      <div className="skeleton-body">
        <div className="skeleton-line shimmer" style={{ width: '90%', height: '20px' }} />
        <div className="skeleton-line shimmer" style={{ width: '100%', height: '14px' }} />
        <div className="skeleton-line shimmer" style={{ width: '85%', height: '14px' }} />
        <div className="skeleton-line shimmer" style={{ width: '60%', height: '14px' }} />
        <div className="skeleton-line shimmer" style={{ width: '40%', height: '12px' }} />
      </div>
    </div>
  )
}

export default function NewsFeed({ articles, loading, error, hasMore, onLoadMore, onRefresh }) {
  const loaderRef = useRef(null)

  useEffect(() => {
    const observer = new IntersectionObserver(
      entries => {
        if (entries[0].isIntersecting && hasMore && !loading) {
          onLoadMore()
        }
      },
      { threshold: 0.1 }
    )
    if (loaderRef.current) observer.observe(loaderRef.current)
    return () => observer.disconnect()
  }, [hasMore, loading, onLoadMore])

  if (error) {
    return (
      <div className="error-state">
        <div className="error-icon">⚠️</div>
        <p className="error-msg">{error}</p>
        <button className="retry-btn" onClick={onRefresh}>Try Again</button>
      </div>
    )
  }

  if (loading && articles.length === 0) {
    return (
      <div className="feed">
        {Array.from({ length: 4 }).map((_, i) => <Skeleton key={i} />)}
      </div>
    )
  }

  if (!loading && articles.length === 0) {
    return (
      <div className="empty-state">
        <div className="empty-icon">📭</div>
        <p>No news found for this category</p>
        <button className="retry-btn" onClick={onRefresh}>Refresh</button>
      </div>
    )
  }

  return (
    <div className="feed">
      {articles.map(article => (
        <NewsCard key={article.id} article={article} />
      ))}

      {loading && articles.length > 0 && (
        <div className="load-more-spinner">
          <div className="spinner" />
        </div>
      )}

      {hasMore && !loading && (
        <div ref={loaderRef} className="loader-trigger" />
      )}

      {!hasMore && articles.length > 0 && (
        <div className="end-message">You're all caught up! ✓</div>
      )}
    </div>
  )
}
