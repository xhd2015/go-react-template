import { BrowserRouter, Route, Routes } from 'react-router-dom';
import { App } from './App';
import { Layout } from './Layout';

export default () => {
    return <BrowserRouter>
        <Routes>
            <Route path="/" element={<Layout />}>
                <Route path="chess" element={<App />}></Route>

                <Route
                    path="*"
                    element={
                        <main style={{ padding: '1rem' }}>
                            <p>Not found!</p>
                        </main>
                    }
                />
            </Route>
        </Routes>
    </BrowserRouter >
}


