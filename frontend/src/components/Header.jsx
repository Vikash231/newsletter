import React from 'react'
import './Header.css'

export default function Header({ onRefresh }) {
  return (
    <header className="header">
      <div className="header-inner">
        <div className="logo">
          <span className="logo-icon">⚡</span>
          <span className="logo-text">InShorts</span>
          <span className="logo-badge">Clone</span>
        </div>
        <div className="header-actions">
          <button className="icon-btn" onClick={onRefresh} title="Refresh">
            ↻
          </button>
        </div>
      </div>
    </header>
  )
}
