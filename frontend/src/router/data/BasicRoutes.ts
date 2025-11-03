import { lazy } from 'react'

import Login from '../../pages/Login/Login'
const Home = lazy(() => import('../../pages/Home'))

type BasicRouteProps = {
    path: string
    Component: React.ComponentType
    isAuthenticated: boolean
}

export const routes: BasicRouteProps[] = [
    {
        path: '/login',
        Component: Login,
        isAuthenticated: false,
    },
    {
        path: '/',
        Component: Home,
        isAuthenticated: true,
    },
]
