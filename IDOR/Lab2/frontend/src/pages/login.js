import React, { useState } from 'react';

const dummyUsers = [
    { id: 1, username: "solninja", password: "pass1", publicKey: "5khiXTUqxMb4hfF1gwVTzSZejnWyfjf32Evai9GBvYUX" },
    { id: 2, username: "solwave", password: "pass2", publicKey: "74HhKaQZp94oEbWgU51qJsSDTLRahVJFaJLqJN5jfs4U" },
    { id: 3, username: "solbird", password: "pass3", publicKey: "22222222222222" },
    { id: 4, username: "solcore", password: "pass4", publicKey: "33333333333333" },
];

const Login = ({ onLogin }) => {
    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');

    const handleLogin = () => {
        const foundUser = dummyUsers.find(
            user => user.username === username && user.password === password
        );

        if (foundUser) {
            localStorage.setItem('user', JSON.stringify(foundUser));
            onLogin(foundUser);
        } else {
            alert("The username or password is incorrect.");
        }
    };

    return (
        <div style={{ display: 'flex', flexDirection: 'column', gap: '10px', maxWidth: '300px', margin: '0 auto' }}>
            <h2>Giriş Yap</h2>
            <input
                value={username}
                onChange={e => setUsername(e.target.value)}
                placeholder="Username"
            />
            <input
                type="password"
                value={password}
                onChange={e => setPassword(e.target.value)}
                placeholder="Password"
            />
            <button onClick={handleLogin}>Giriş</button>
        </div>
    );
};

export default Login;
