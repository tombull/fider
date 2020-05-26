import "./Error.page.scss";

import React from "react";
import { TenantLogo } from "@teamdream/components";
import { useTeamdream } from "@teamdream/hooks";

interface ErrorPageProps {
  error: Error;
  errorInfo: React.ErrorInfo;
  showDetails?: boolean;
}

export const ErrorPage = (props: ErrorPageProps) => {
  const teamdream = useTeamdream();

  return (
    <div id="p-error" className="container failure-page">
      <TenantLogo size={100} useTeamdreamIfEmpty={true} />
      <h1>Shoot! Well, this is unexpected…</h1>
      <p>An error has occurred and we're working to fix the problem!</p>
      {teamdream.settings && (
        <span>
          Take me back to <a href={teamdream.settings.baseURL}>{teamdream.settings.baseURL}</a> home page.
        </span>
      )}
      {props.showDetails && (
        <pre className="error">
          {props.error.toString()}
          {props.errorInfo.componentStack}
        </pre>
      )}
    </div>
  );
};
