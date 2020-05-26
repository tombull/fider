import "./Home.page.scss";

import React, { useState } from "react";
import { Post, Tag, PostStatus } from "@teamdream/models";
import { MultiLineText, Hint } from "@teamdream/components";
import { SimilarPosts } from "./components/SimilarPosts";
import { FaRegLightbulb } from "react-icons/fa";
import { PostInput } from "./components/PostInput";
import { PostsContainer } from "./components/PostsContainer";
import { useTeamdream } from "@teamdream/hooks";

export interface HomePageProps {
  posts: Post[];
  tags: Tag[];
  countPerStatus: { [key: string]: number };
}

export interface HomePageState {
  title: string;
}

const Lonely = () => {
  const teamdream = useTeamdream();

  return (
    <div className="l-lonely center">
      <Hint
        permanentCloseKey="at-least-3-posts"
        condition={teamdream.session.isAuthenticated && teamdream.session.user.isAdministrator}
      >
        It's recommended that you post <strong>at least 3</strong> suggestions here before sharing this site. The
        initial content is key to start the interactions with your audience.
      </Hint>
      <p>
        <FaRegLightbulb />
      </p>
      <p>It's lonely out here. Start by sharing a suggestion!</p>
    </div>
  );
};

const defaultWelcomeMessage = `We'd love to hear what you're thinking about.

What can we do better? This is the place for you to vote, discuss and share ideas.`;

const HomePage = (props: HomePageProps) => {
  const teamdream = useTeamdream();
  const [title, setTitle] = useState("");

  const isLonely = () => {
    const len = Object.keys(props.countPerStatus).length;
    if (len === 0) {
      return true;
    }

    if (len === 1 && PostStatus.Deleted.value in props.countPerStatus) {
      return true;
    }

    return false;
  };

  return (
    <div id="p-home" className="page container">
      <div className="row">
        <div className="l-welcome-col col-md-4">
          <MultiLineText
            className="welcome-message"
            text={teamdream.session.tenant.welcomeMessage || defaultWelcomeMessage}
            style="full"
          />
          <PostInput
            placeholder={teamdream.session.tenant.invitation || "Enter your suggestion here..."}
            onTitleChanged={setTitle}
          />
        </div>
        <div className="l-posts-col col-md-8">
          {isLonely() ? (
            <Lonely />
          ) : title ? (
            <SimilarPosts title={title} tags={props.tags} />
          ) : (
            <PostsContainer posts={props.posts} tags={props.tags} countPerStatus={props.countPerStatus} />
          )}
        </div>
      </div>
    </div>
  );
};

export default HomePage;
