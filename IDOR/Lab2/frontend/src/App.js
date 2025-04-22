import React, { useEffect, useState } from 'react';
import Login from './pages/login';
import Dashboard from './pages/dashboard';

// Developed by Çetin Boran Mesüm

function App() {
  const [user, setUser] = useState(null);

  useEffect(() => {
    const storedUser = localStorage.getItem('user');
    if (storedUser) {
      setUser(JSON.parse(storedUser));
    }
  }, []);

  const handleLogin = (userData) => {
    setUser(userData);
  };

  return (
    <div className="App">
      {user ? (
        <Dashboard user={user} />
      ) : (
        <Login onLogin={handleLogin} />
      )}
    </div>
  );
}

export default App;
