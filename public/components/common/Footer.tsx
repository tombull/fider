import "./Footer.scss";

import React from "react";
import { PrivacyPolicy, TermsOfService } from "@teamdream/components";
import { useTeamdream } from "@teamdream/hooks";

export const Footer = () => {
  const teamdream = useTeamdream();

  return (
    <div id="c-footer">
      <div className="container">
        {teamdream.settings.hasLegal && (
          <div className="l-links">
            <PrivacyPolicy />
            &middot;
            <TermsOfService />
          </div>
        )}
        <a className="l-powered" target="_blank" href="https://teamdream.co.uk/">
          <img src="https://teamdream.co.uk/images/TeamDream-Logo.svg" alt="Teamdream" />
          <span>Powered by Teamdream</span>
        </a>
      </div>
    </div>
  );
};
