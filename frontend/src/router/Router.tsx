import { Route, BrowserRouter as Router, Routes as Switch } from "react-router";
import { routes } from "./Routes";

export default function AppRouter() {
    return (
        <Router>
            <Switch>
                {routes.map((route, index) => {
                    return <Route key={index} {...route} />;
                })}
            </Switch>
        </Router>
    );
}
