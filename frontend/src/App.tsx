import './App.css'

function App() {
  return (
    <main className="app-shell">
      <header className="app-header">
        <div>
          <p className="eyebrow">Self-hosted task manager</p>
          <h1>MaxTasks</h1>
        </div>
        <span className="status-pill">Foundation</span>
      </header>

      <section className="welcome-card" aria-labelledby="welcome-title">
        <p className="card-label">Workspace ready</p>
        <h2 id="welcome-title">A calmer way to organize your work.</h2>
        <p>
          The application shell is ready. Task lists, authentication, and the
          API are the next milestones.
        </p>
        <button type="button" disabled>
          Create your first list
        </button>
      </section>

      <footer>React · Go · PostgreSQL · Docker</footer>
    </main>
  )
}

export default App
