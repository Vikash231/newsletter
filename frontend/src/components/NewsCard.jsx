import React, { useState } from 'react'
import './NewsCard.css'

function formatTime(dateStr) {
  const date = new Date(dateStr)
  const now = new Date()
  const diff = Math.floor((now - date) / 1000)
  if (diff < 60) return 'just now'
  if (diff < 3600) return `${Math.floor(diff / 60)}m ago`
  if (diff < 86400) return `${Math.floor(diff / 3600)}h ago`
  return `${Math.floor(diff / 86400)}d ago`
}

function wordCount(text) {
  return text.trim().split(/\s+/).length
}

export default function NewsCard({ article }) {
  const [bookmarked, setBookmarked] = useState(false)
  const [shared, setShared] = useState(false)
  const [imgError, setImgError] = useState(false)

  const handleShare = async () => {
    if (navigator.share) {
      try {
        await navigator.share({ title: article.title, text: article.content, url: article.source_url })
      } catch (_) {}
    } else {
      await navigator.clipboard.writeText(article.source_url)
      setShared(true)
      setTimeout(() => setShared(false), 2000)
    }
  }

  const handleBookmark = () => setBookmarked(b => !b)

  return (
    <article className="news-card">
      {!imgError && article.image_url && (
        <div className="card-image">
          <img
            src={article.image_url}
            alt={article.title}
            loading="lazy"
            onError={() => setImgError(true)}
          />
          <span className={`category-pill ${article.category}`}>
            {article.category}
          </span>
        </div>
      )}

      <div className="card-body">
        <h2 className="card-title">{article.title}</h2>

        <p className="card-content">{article.content}</p>

        <div className="word-count">
          {wordCount(article.content)} words
        </div>

        <div className="card-meta">
          <div className="meta-left">
            <span className="source">{article.source}</span>
            <span className="dot">·</span>
            <span className="time">{formatTime(article.published_at)}</span>
            <span className="dot">·</span>
            <span className="read-time">{article.read_time}s read</span>
          </div>
        </div>

        <div className="card-actions">
          <a
            href={article.source_url}
            target="_blank"
            rel="noopener noreferrer"
            className="action-link"
          >
            Read Full Story →
          </a>
          <div className="action-buttons">
            <button
              className={`action-btn ${shared ? 'active' : ''}`}
              onClick={handleShare}
              title="Share"
            >
              {shared ? '✓ Copied' : '↗ Share'}
            </button>
            <button
              className={`action-btn bookmark-btn ${bookmarked ? 'active' : ''}`}
              onClick={handleBookmark}
              title={bookmarked ? 'Remove bookmark' : 'Bookmark'}
            >
              {bookmarked ? '🔖 Saved' : '☆ Save'}
            </button>
          </div>
        </div>
      </div>
    </article>
  )
}
