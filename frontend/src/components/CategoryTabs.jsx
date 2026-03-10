import React from 'react'
import './CategoryTabs.css'

const CATEGORY_ICONS = {
  all: '🗞️',
  technology: '💻',
  sports: '⚽',
  business: '📈',
  science: '🔬',
  entertainment: '🎬',
  health: '🏥',
  politics: '🏛️',
  world: '🌍',
}

export default function CategoryTabs({ categories, active, onChange }) {
  return (
    <div className="categories-wrapper">
      <div className="categories">
        {categories.map(cat => (
          <button
            key={cat}
            className={`category-tab ${active === cat ? 'active' : ''}`}
            onClick={() => onChange(cat)}
          >
            <span className="cat-icon">{CATEGORY_ICONS[cat] || '📰'}</span>
            <span className="cat-label">{cat.charAt(0).toUpperCase() + cat.slice(1)}</span>
          </button>
        ))}
      </div>
    </div>
  )
}
