import { lazy, ComponentType } from "react";

type LazyImport = () => Promise<{ default: ComponentType<any> }>;

const MAX_RETRIES = 10;
const INTERVAL = 500;

const retry = (fn: LazyImport, retriesLeft = MAX_RETRIES): Promise<{ default: ComponentType<any> }> => {
  return new Promise((resolve, reject) => {
    fn()
      .then(resolve)
      .catch((err) => {
        setTimeout(() => {
          if (retriesLeft === 1) {
            reject(new Error(`${err} after ${MAX_RETRIES} retries`));
            return;
          }
          retry(fn, retriesLeft - 1).then(resolve, reject);
        }, INTERVAL);
      });
  });
};

const load = (fn: LazyImport) => lazy(() => retry(() => fn()));

export const AsyncHomePage = load(() =>
  import(
    /* webpackChunkName: "Home.page" */
    "@teamdream/pages/Home/Home.page"
  )
);

export const AsyncShowPostPage = load(() =>
  import(
    /* webpackChunkName: "ShowPost.page" */
    "@teamdream/pages/ShowPost/ShowPost.page"
  )
);

export const AsyncManageMembersPage = load(() =>
  import(
    /* webpackChunkName: "ManageMembers.page" */
    "@teamdream/pages/Administration/pages/ManageMembers.page"
  )
);

export const AsyncManageTagsPage = load(() =>
  import(
    /* webpackChunkName: "ManageTags.page" */
    "@teamdream/pages/Administration/pages/ManageTags.page"
  )
);

export const AsyncPrivacySettingsPage = load(() =>
  import(
    /* webpackChunkName: "PrivacySettings.page" */
    "@teamdream/pages/Administration/pages/PrivacySettings.page"
  )
);

export const AsyncExportPage = load(() =>
  import(
    /* webpackChunkName: "Export.page" */
    "@teamdream/pages/Administration/pages/Export.page"
  )
);

export const AsyncInvitationsPage = load(() =>
  import(
    /* webpackChunkName: "Invitations.page" */
    "@teamdream/pages/Administration/pages/Invitations.page"
  )
);

export const AsyncManageAuthenticationPage = load(() =>
  import(
    /* webpackChunkName: "ManageAuthentication.page" */
    "@teamdream/pages/Administration/pages/ManageAuthentication.page"
  )
);

export const AsyncAdvancedSettingsPage = load(() =>
  import(
    /* webpackChunkName: "AdvancedSettings.page" */
    "@teamdream/pages/Administration/pages/AdvancedSettings.page"
  )
);

export const AsyncGeneralSettingsPage = load(() =>
  import(
    /* webpackChunkName: "GeneralSettings.page" */
    "@teamdream/pages/Administration/pages/GeneralSettings.page"
  )
);

export const AsyncSignInPage = load(() =>
  import(
    /* webpackChunkName: "SignIn.page" */
    "@teamdream/pages/SignIn/SignIn.page"
  )
);

export const AsyncSignUpPage = load(() =>
  import(
    /* webpackChunkName: "SignUp.page" */
    "@teamdream/pages/SignUp/SignUp.page"
  )
);

export const AsyncCompleteSignInProfilePage = load(() =>
  import(
    /* webpackChunkName: "CompleteSignInProfile.page" */
    "@teamdream/pages/CompleteSignInProfile/CompleteSignInProfile.page"
  )
);

export const AsyncMyNotificationsPage = load(() =>
  import(
    /* webpackChunkName: "MyNotifications.page" */
    "@teamdream/pages/MyNotifications/MyNotifications.page"
  )
);

export const AsyncMySettingsPage = load(() =>
  import(
    /* webpackChunkName: "MySettings.page" */
    "@teamdream/pages/MySettings/MySettings.page"
  )
);

export const AsyncBillingPage = load(() =>
  import(
    /* webpackChunkName: "Billing.page" */
    "@teamdream/pages/Administration/pages/Billing.page"
  )
);

export const AsyncOAuthEchoPage = load(() =>
  import(
    /* webpackChunkName: "OAuthEcho.page" */
    "@teamdream/pages/OAuthEcho/OAuthEcho.page"
  )
);

export const AsyncUIToolkitPage = load(() =>
  import(
    /* webpackChunkName: "UIToolkit.page" */
    "@teamdream/pages/UI/UIToolkit.page"
  )
);
