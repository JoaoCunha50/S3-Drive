import { lazy } from "react";
import { type RouteProps } from "react-router";

import Login from "../pages/Login";
const Home = lazy(() => import("../pages/Home"));

export const routes: RouteProps[] = [
    {
        path: "/login",
        Component: Login,
    },
    {
        path: "/",
        Component: Home,
    },
];
