import * as Pages from "@teamdream/AsyncPages";

interface PageConfiguration {
  regex: RegExp;
  component: any;
  showHeader: boolean;
}

const route = (path: string, component: any, showHeader: boolean = true): PageConfiguration => {
  path = path.replace("/", "/").replace(":number", "\\d+").replace(":string", ".+").replace("*", "/?.*");

  const regex = new RegExp(`^${path}$`);
  return { regex, component, showHeader };
};

const pathRegex = [
  route("", Pages.AsyncHomePage),
  route("/posts/:number*", Pages.AsyncShowPostPage),
  route("/admin/members", Pages.AsyncManageMembersPage),
  route("/admin/tags", Pages.AsyncManageTagsPage),
  route("/admin/privacy", Pages.AsyncPrivacySettingsPage),
  route("/admin/billing", Pages.AsyncBillingPage),
  route("/admin/export", Pages.AsyncExportPage),
  route("/admin/invitations", Pages.AsyncInvitationsPage),
  route("/admin/authentication", Pages.AsyncManageAuthenticationPage),
  route("/admin/advanced", Pages.AsyncAdvancedSettingsPage),
  route("/admin", Pages.AsyncGeneralSettingsPage),
  route("/signin", Pages.AsyncSignInPage, false),
  route("/signup", Pages.AsyncSignUpPage, false),
  route("/signin/verify", Pages.AsyncCompleteSignInProfilePage),
  route("/invite/verify", Pages.AsyncCompleteSignInProfilePage),
  route("/notifications", Pages.AsyncMyNotificationsPage),
  route("/settings", Pages.AsyncMySettingsPage),
  route("/oauth/:string/echo", Pages.AsyncOAuthEchoPage, false),
  route("/-/ui", Pages.AsyncUIToolkitPage),
];

export const resolveRootComponent = (path: string): PageConfiguration => {
  if (path.length > 0 && path.charAt(path.length - 1) === "/") {
    path = path.substring(0, path.length - 1);
  }
  for (const entry of pathRegex) {
    if (entry && entry.regex.test(path)) {
      return entry;
    }
  }
  throw new Error(`Component not found for route ${path}.`);
};
