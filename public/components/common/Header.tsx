import "./Header.scss";

import React, { useState, useEffect } from "react";
import { SignInModal, EnvironmentInfo, Avatar, TenantLogo, TenantStatusInfo } from "@teamdream/components";
import { actions } from "@teamdream/services";
import { FaUser, FaCog, FaCaretDown } from "react-icons/fa";
import { useTeamdream } from "@teamdream/hooks";

export const Header = () => {
  const teamdream = useTeamdream();
  const [isSignInModalOpen, setIsSignInModalOpen] = useState(false);
  const [unreadNotifications, setUnreadNotifications] = useState(0);

  useEffect(() => {
    if (teamdream.session.isAuthenticated) {
      actions.getTotalUnreadNotifications().then((result) => {
        if (result.ok && result.data > 0) {
          setUnreadNotifications(result.data);
        }
      });
    }
  }, [teamdream.session.isAuthenticated]);

  const showModal = () => {
    if (!teamdream.session.isAuthenticated) {
      setIsSignInModalOpen(true);
    }
  };

  const hideModal = () => setIsSignInModalOpen(false);

  const items = teamdream.session.isAuthenticated && (
    <div className="c-menu-user">
      <div className="c-menu-user-heading">
        <FaUser /> <span>{teamdream.session.user.name}</span>
      </div>
      <a href="/settings" className="c-menu-user-item">
        Settings
      </a>
      <a href="/notifications" className="c-menu-user-item">
        Notifications
        {unreadNotifications > 0 && <div className="c-unread-count">{unreadNotifications}</div>}
      </a>
      <div className="c-menu-user-divider" />
      {teamdream.session.user.isCollaborator && [
        <div key={1} className="c-menu-user-heading">
          <FaCog /> <span>Administration</span>
        </div>,
        <a key={2} href="/admin" className="c-menu-user-item">
          Site Settings
        </a>,
        <div key={5} className="c-menu-user-divider" />,
      ]}
      <a href="/signout?redirect=/" className="c-menu-user-item signout">
        Sign out
      </a>
    </div>
  );

  const showRightMenu = teamdream.session.isAuthenticated || !teamdream.session.tenant.isPrivate;
  return (
    <div id="c-header">
      <EnvironmentInfo />
      <SignInModal isOpen={isSignInModalOpen} onClose={hideModal} />
      <div className="c-menu">
        <div className="container">
          <a href="/" className="c-menu-item-title">
            <TenantLogo size={100} />
            <span>{teamdream.session.tenant.name}</span>
          </a>
          {showRightMenu && (
            <div onClick={showModal} className="c-menu-item-signin">
              {teamdream.session.isAuthenticated && <Avatar user={teamdream.session.user} />}
              {unreadNotifications > 0 && <div className="c-unread-dot" />}
              {!teamdream.session.isAuthenticated && <span>Sign in</span>}
              {teamdream.session.isAuthenticated && <FaCaretDown />}
              {items}
            </div>
          )}
        </div>
      </div>
      <TenantStatusInfo />
    </div>
  );
};
