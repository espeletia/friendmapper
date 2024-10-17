import { Link, useRouteError } from "react-router-dom";

const ErrorRoute = () => {
  const error = useRouteError() as { statusText: string; message: string };

  return (
    <div>
      <h1>Oops!</h1>
      <p>Něco se pokazilo.</p>
      <p>
        <i>{error.statusText || error.message}</i>
      </p>
      <Link to="/">Zpět na domovskou stránku</Link>
    </div>
  );
};

export default ErrorRoute;
