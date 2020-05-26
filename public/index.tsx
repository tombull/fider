import React, { Suspense } from "react";
import ReactDOM from "react-dom";
import { resolveRootComponent } from "@teamdream/router";
import { Header, Footer, Loader } from "@teamdream/components/common";
import { ErrorBoundary } from "@teamdream/components";
import { classSet, Teamdream, TeamdreamContext, actions, navigator } from "@teamdream/services";
import { IconContext } from "react-icons";

const Loading = () => (
  <div className="page">
    <Loader />
  </div>
);

import "@teamdream/assets/styles/index.scss";

const logProductionError = (err: Error) => {
  if (Teamdream.isProduction()) {
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
  let teamdream;

  if (!navigator.isBrowserSupported()) {
    navigator.goTo("/browser-not-supported");
    return;
  }

  teamdream = Teamdream.initialize();

  __webpack_nonce__ = teamdream.session.contextID;
  __webpack_public_path__ = `${teamdream.settings.globalAssetsURL}/assets/`;

  const config = resolveRootComponent(location.pathname);
  document.body.className = classSet({
    "is-authenticated": teamdream.session.isAuthenticated,
    "is-staff": teamdream.session.isAuthenticated && teamdream.session.user.isCollaborator,
  });
  ReactDOM.render(
    <React.StrictMode>
      <ErrorBoundary onError={logProductionError}>
        <TeamdreamContext.Provider value={teamdream}>
          <IconContext.Provider value={{ className: "icon" }}>
            {config.showHeader && <Header />}
            <Suspense fallback={<Loading />}>{React.createElement(config.component, teamdream.session.props)}</Suspense>
            {config.showHeader && <Footer />}
          </IconContext.Provider>
        </TeamdreamContext.Provider>
      </ErrorBoundary>
    </React.StrictMode>,
    document.getElementById("root")
  );
})();
