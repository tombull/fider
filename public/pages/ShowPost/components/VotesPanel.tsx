import "./VotesPanel.scss";

import React, { useState } from "react";
import { Post, Vote } from "@teamdream/models";
import { Avatar } from "@teamdream/components";
import { Teamdream, classSet } from "@teamdream/services";
import { useTeamdream } from "@teamdream/hooks";
import { VotesModal } from "./VotesModal";

interface VotesPanelProps {
  post: Post;
  votes: Vote[];
}

export const VotesPanel = (props: VotesPanelProps) => {
  const teamdream = useTeamdream();
  const [isVotesModalOpen, setIsVotesModalOpen] = useState(false);

  const openModal = () => {
    if (canShowAll()) {
      setIsVotesModalOpen(true);
    }
  };

  const closeModal = () => setIsVotesModalOpen(false);
  const canShowAll = () => teamdream.session.isAuthenticated && Teamdream.session.user.isCollaborator;

  const extraVotesCount = props.post.votesCount - props.votes.length;
  const moreVotesClassName = classSet({
    "l-votes-more": true,
    clickable: canShowAll(),
  });

  return (
    <>
      <VotesModal post={props.post} isOpen={isVotesModalOpen} onClose={closeModal} />
      <span className="subtitle">Voters</span>
      <div className="l-votes-list">
        {props.votes.map((x) => (
          <Avatar key={x.user.id} user={x.user} />
        ))}
        {extraVotesCount > 0 && (
          <span onClick={openModal} className={moreVotesClassName}>
            +{extraVotesCount} more
          </span>
        )}
        {props.votes.length > 0 && extraVotesCount === 0 && canShowAll() && (
          <span onClick={openModal} className={moreVotesClassName}>
            see details
          </span>
        )}
        {props.votes.length === 0 && <span className="info">None yet</span>}
      </div>
    </>
  );
};
