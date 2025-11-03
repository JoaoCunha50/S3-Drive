import { Suspense } from 'react'
import { Route, BrowserRouter as Router, Routes as Switch } from 'react-router'
import LoadingPage from '../pages/LoadingPage/LoadingPage'
import { routes } from './data/BasicRoutes'

export default function AppRouter() {
    return (
        <Suspense fallback={<LoadingPage />}>
            <Router>
                <Switch>
                    {routes.map((route, index) => {
                        return <Route key={index} {...route} />
                    })}
                </Switch>
            </Router>
        </Suspense>
    )
}
