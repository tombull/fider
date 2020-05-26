import "./GeneralSettings.page.scss";

import React from "react";

import { Button, ButtonClickEvent, TextArea, Form, Input, ImageUploader } from "@teamdream/components/common";
import { actions, Failure, Teamdream } from "@teamdream/services";
import { FaCogs } from "react-icons/fa";
import { AdminBasePage } from "../components/AdminBasePage";
import { ImageUpload } from "@teamdream/models";

interface GeneralSettingsPageState {
  logo?: ImageUpload;
  title: string;
  invitation: string;
  welcomeMessage: string;
  cname: string;
  error?: Failure;
}

export default class GeneralSettingsPage extends AdminBasePage<{}, GeneralSettingsPageState> {
  public id = "p-admin-general";
  public name = "general";
  public icon = FaCogs;
  public title = "General";
  public subtitle = "Manage your site settings";

  constructor(props: {}) {
    super(props);

    this.state = {
      title: Teamdream.session.tenant.name,
      cname: Teamdream.session.tenant.cname,
      welcomeMessage: Teamdream.session.tenant.welcomeMessage,
      invitation: Teamdream.session.tenant.invitation,
    };
  }

  private handleSave = async (e: ButtonClickEvent) => {
    const result = await actions.updateTenantSettings(this.state);
    if (result.ok) {
      e.preventEnable();
      location.href = `/`;
    } else if (result.error) {
      this.setState({ error: result.error });
    }
  };

  public dnsInstructions(): JSX.Element {
    const isApex = this.state.cname.split(".").length <= 2;
    const recordType = isApex ? "ALIAS" : "CNAME";
    return (
      <>
        <strong>{this.state.cname}</strong> {recordType}{" "}
        <strong>
          {Teamdream.session.tenant.subdomain}
          {Teamdream.settings.domain}
        </strong>
      </>
    );
  }

  private setTitle = (title: string): void => {
    this.setState({ title });
  };

  private setWelcomeMessage = (welcomeMessage: string): void => {
    this.setState({ welcomeMessage });
  };

  private setInvitation = (invitation: string): void => {
    this.setState({ invitation });
  };

  private setLogo = (logo: ImageUpload): void => {
    this.setState({ logo });
  };

  private setCNAME = (cname: string): void => {
    this.setState({ cname });
  };

  public content() {
    return (
      <Form error={this.state.error}>
        <Input
          field="title"
          label="Title"
          maxLength={60}
          value={this.state.title}
          disabled={!Teamdream.session.user.isAdministrator}
          onChange={this.setTitle}
        >
          <p className="info">
            The title is used on the header, emails, notifications and SEO content. Keep it short and simple. The
            product/service name is usually the best choice.
          </p>
        </Input>

        <TextArea
          field="welcomeMessage"
          label="Welcome Message"
          value={this.state.welcomeMessage}
          disabled={!Teamdream.session.user.isAdministrator}
          onChange={this.setWelcomeMessage}
        >
          <p className="info">
            The message is shown on this site's home page. Use it to help visitors understad what this space is about
            and the importance of their feedback. Leave it empty for a default message.
          </p>
        </TextArea>

        <Input
          field="invitation"
          label="Invitation"
          maxLength={60}
          value={this.state.invitation}
          disabled={!Teamdream.session.user.isAdministrator}
          placeholder="Enter your suggestion here..."
          onChange={this.setInvitation}
        >
          <p className="info">
            This text is used as a placeholder for the suggestion's text box. Use it to invite your visitors into
            sharing their suggestions and feedback. Leave it empty for a default message.
          </p>
        </Input>

        <ImageUploader
          label="Logo"
          field="logo"
          bkey={Teamdream.session.tenant.logoBlobKey}
          previewMaxWidth={200}
          disabled={!Teamdream.session.user.isAdministrator}
          onChange={this.setLogo}
        >
          <p className="info">
            We accept JPG, GIF and PNG images, smaller than 100KB and with an aspect ratio of 1:1 with minimum
            dimensions of 200x200 pixels.
          </p>
        </ImageUploader>

        {!Teamdream.isSingleHostMode() && (
          <Input
            field="cname"
            label="Custom Domain"
            maxLength={100}
            placeholder="feedback.yourcompany.com"
            value={this.state.cname}
            disabled={!Teamdream.session.user.isAdministrator}
            onChange={this.setCNAME}
          >
            <div className="info">
              {this.state.cname ? (
                [
                  <p key={0}>Enter the following record into your DNS zone records:</p>,
                  <p key={1}>{this.dnsInstructions()}</p>,
                  <p key={2}>
                    Please note that it may take up to 72 hours for the change to take effect worldwide due to DNS
                    propagation.
                  </p>,
                ]
              ) : (
                <p>
                  Custom domains allow you to access your app via your own domain name (for example,{" "}
                  <code>feedback.yourcompany.com</code>
                  ).
                </p>
              )}
            </div>
          </Input>
        )}

        <div className="field">
          <Button disabled={!Teamdream.session.user.isAdministrator} color="positive" onClick={this.handleSave}>
            Save
          </Button>
        </div>
      </Form>
    );
  }
}
