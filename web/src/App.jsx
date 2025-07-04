import React, { useState } from 'react';
import './index.css';

export default function App() {
    const [query, setQuery] = useState('select * from example;');
    const [result, setResult] = useState('');
    const [loading, setLoading] = useState(false);

    const handleSubmit = async () => {
        if (!query.trim()) return;
        setLoading(true);
        setResult('');
        try {
            const response = await fetch('http://127.0.0.1:21101/zeus/v1/soar', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({
                    soar: {
                        query,
                    },
                }),
            });
            const data = await response.json();
            setResult(data.data);
        } catch (error) {
            setResult(`Error: ${error.message}`);
        } finally {
            setLoading(false);
        }
    };

    return (
        <div className="container">
            <h1 className="title">Soar SQL 测试工具</h1>
            <textarea
                value={query}
                onChange={(e) => setQuery(e.target.value)}
                rows={8}
                className="textarea"
                placeholder="请输入 SQL 查询语句..."
            />
            <button onClick={handleSubmit} className="submit-button" disabled={loading}>
                {loading ? '执行中...' : '执行分析'}
            </button>
            {result && (
                <div className="result-box">
                    <h2>分析结果</h2>
                    <pre className="result">{result}</pre>
                </div>
            )}
        </div>
    );
}
