import { createBrowserRouter, RouterProvider } from "react-router-dom";
import HomeRoute from "./routes/home";
import ErrorRoute from "./routes/error";
import { LeafletProvider } from "./components/Leaflet/context/LeafletContext";
import MapRoute from "./routes/map";
import Root from "./routes/root";
import AuthDialogRoute from "./routes/authDialog";
import SearchOverlayRoute from "./routes/searchOverlay";

const router = createBrowserRouter([
  {
    path: "/",
    element: <Root />,
    errorElement: <ErrorRoute />,
    children: [
      {
        path: "/",
        element: <HomeRoute />,
        errorElement: <ErrorRoute />,
      },
      {
        path: "map",
        element: <MapRoute />,
        errorElement: <ErrorRoute />,
      },
      {
        path: "auth",
        element: <AuthDialogRoute />,
        errorElement: <ErrorRoute />,
      },
      {
        path: "search",
        element: <SearchOverlayRoute />,
        errorElement: <ErrorRoute />,
      }
    ],
  },
]);

function App() {
  return (
    <LeafletProvider>
      <RouterProvider router={router} />
    </LeafletProvider>
  );
}

export default App;
