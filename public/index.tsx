import React, { Suspense } from "react";
import ReactDOM from "react-dom";
import { resolveRootComponent } from "@fider/router";
import { Header, Footer, Loader } from "@fider/components/common";
import { ErrorBoundary } from "@fider/components";
import { classSet, Fider, FiderContext, actions, navigator } from "@fider/services";
import { IconContext } from "react-icons";

const Loading = () => (
  <div className="page">
    <Loader />
  </div>
);

import "@fider/assets/styles/index.scss";

const logProductionError = (err: Error) => {
  if (Fider.isProduction()) {
    console.error(err); // tslint:disable-line
    actions.logError(`react.ErrorBoundary: ${err.message}`, err);
  }
};

window.addEventListener("unhandledrejection", (evt: PromiseRejectionEvent) => {
  if (evt.reason instanceof Error) {
    actions.logError(`window.unhandledrejection: ${evt.reason.message}`, evt.reason);
  } else if (evt.reason) {
    actions.logError(`window.unhandledrejection: ${evt.reason.toString()}`);
  }
});

window.addEventListener("error", (evt: ErrorEvent) => {
  if (evt.error && evt.colno > 0 && evt.lineno > 0) {
    actions.logError(`window.error: ${evt.message}`, evt.error);
  }
});

(() => {
  let fider;

  if (!navigator.isBrowserSupported()) {
    navigator.goTo("/browser-not-supported");
    return;
  }

  fider = Fider.initialize();

  __webpack_nonce__ = fider.session.contextID;
  __webpack_public_path__ = `${fider.settings.globalAssetsURL}/assets/`;

  const config = resolveRootComponent(location.pathname);
  document.body.className = classSet({
    "is-authenticated": fider.session.isAuthenticated,
    "is-staff": fider.session.isAuthenticated && fider.session.user.isCollaborator,
  });
  ReactDOM.render(
    <React.StrictMode>
      <ErrorBoundary onError={logProductionError}>
        <FiderContext.Provider value={fider}>
          <IconContext.Provider value={{ className: "icon" }}>
            {config.showHeader && <Header />}
            <Suspense fallback={<Loading />}>{React.createElement(config.component, fider.session.props)}</Suspense>
            {config.showHeader && <Footer />}
          </IconContext.Provider>
        </FiderContext.Provider>
      </ErrorBoundary>
    </React.StrictMode>,
    document.getElementById("root")
  );
})();
