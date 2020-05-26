import React, { useState } from "react";
import { PostStatus, Post } from "@teamdream/models";
import { actions, navigator, Failure } from "@teamdream/services";
import { Form, Modal, Button, List, ListItem, TextArea } from "@teamdream/components";
import { useTeamdream } from "@teamdream/hooks";

interface ModerationPanelProps {
  post: Post;
}

export const ModerationPanel = (props: ModerationPanelProps) => {
  const teamdream = useTeamdream();
  const [showConfirmation, setShowConfirmation] = useState(false);
  const [text, setText] = useState("");
  const [error, setError] = useState<Failure>();

  const hideModal = async () => setShowConfirmation(false);
  const showModal = async () => setShowConfirmation(true);

  const handleDelete = async () => {
    const response = await actions.deletePost(props.post.number, text);
    if (response.ok) {
      hideModal();
      navigator.goHome();
    } else if (response.error) {
      setError(response.error);
    }
  };

  const status = PostStatus.Get(props.post.status);
  if (!teamdream.session.isAuthenticated || !teamdream.session.user.isAdministrator || status.closed) {
    return null;
  }

  const modal = (
    <Modal.Window isOpen={showConfirmation} onClose={hideModal} center={false} size="large">
      <Modal.Content>
        <Form error={error}>
          <TextArea
            field="text"
            onChange={setText}
            value={text}
            placeholder="Why are you deleting this post? (optional)"
          >
            <span className="info">
              This operation <strong>cannot</strong> be undone.
            </span>
          </TextArea>
        </Form>
      </Modal.Content>

      <Modal.Footer>
        <Button color="danger" onClick={handleDelete}>
          Delete
        </Button>
        <Button color="cancel" onClick={hideModal}>
          Cancel
        </Button>
      </Modal.Footer>
    </Modal.Window>
  );

  return (
    <>
      {modal}
      <span className="subtitle">Moderation</span>
      <List>
        <ListItem>
          <Button color="danger" size="tiny" fluid={true} onClick={showModal}>
            Delete
          </Button>
        </ListItem>
      </List>
    </>
  );
};
