import React, { useState } from 'react'
import Header from './components/Header.jsx'
import CategoryTabs from './components/CategoryTabs.jsx'
import NewsFeed from './components/NewsFeed.jsx'
import { useNews, useCategories } from './hooks/useNews.js'
import './App.css'

export default function App() {
  const [activeCategory, setActiveCategory] = useState('all')
  const { categories } = useCategories()
  const { articles, loading, error, hasMore, loadMore, refresh } = useNews(activeCategory)

  const handleCategoryChange = (cat) => {
    setActiveCategory(cat)
  }

  return (
    <div className="app">
      <Header onRefresh={refresh} />
      <CategoryTabs
        categories={categories}
        active={activeCategory}
        onChange={handleCategoryChange}
      />
      <main className="main">
        <NewsFeed
          articles={articles}
          loading={loading}
          error={error}
          hasMore={hasMore}
          onLoadMore={loadMore}
          onRefresh={refresh}
        />
      </main>
    </div>
  )
}
