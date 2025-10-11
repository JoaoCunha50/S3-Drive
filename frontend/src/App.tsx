import { LoaderCircle as Spinner } from "lucide-react";
import { Suspense } from "react";
import "./App.css";
import AppRouter from "./router/Router";

function App() {
    return (
        <Suspense fallback={<Spinner className="spinner" />}>
            <AppRouter />
        </Suspense>
    );
}

export default App;
