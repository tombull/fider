import React, { useState, useRef } from "react";

import { Post, ImageUpload } from "@teamdream/models";
import { Avatar, UserName, Button, TextArea, Form, MultiImageUploader } from "@teamdream/components/common";
import { SignInModal } from "@teamdream/components";

import { cache, actions, Failure, Teamdream } from "@teamdream/services";
import { useTeamdream } from "@teamdream/hooks";

interface CommentInputProps {
  post: Post;
}

const CACHE_TITLE_KEY = "CommentInput-Comment-";

export const CommentInput = (props: CommentInputProps) => {
  const getCacheKey = () => `${CACHE_TITLE_KEY}${props.post.id}`;

  const teamdream = useTeamdream();
  const inputRef = useRef<HTMLTextAreaElement>();
  const [content, setContent] = useState((teamdream.session.isAuthenticated && cache.session.get(getCacheKey())) || "");
  const [isSignInModalOpen, setIsSignInModalOpen] = useState(false);
  const [attachments, setAttachments] = useState<ImageUpload[]>([]);
  const [error, setError] = useState<Failure | undefined>(undefined);

  const commentChanged = (newContent: string) => {
    cache.session.set(getCacheKey(), newContent);
    setContent(newContent);
  };

  const hideModal = () => setIsSignInModalOpen(false);
  const clearError = () => setError(undefined);

  const submit = async () => {
    clearError();

    const result = await actions.createComment(props.post.number, content, attachments);
    if (result.ok) {
      cache.session.remove(getCacheKey());
      location.reload();
    } else {
      setError(result.error);
    }
  };

  const handleOnFocus = () => {
    if (!teamdream.session.isAuthenticated && inputRef.current) {
      inputRef.current.blur();
      setIsSignInModalOpen(true);
    }
  };

  return (
    <>
      <SignInModal isOpen={isSignInModalOpen} onClose={hideModal} />
      <div className={`c-comment-input ${Teamdream.session.isAuthenticated && "m-authenticated"}`}>
        {Teamdream.session.isAuthenticated && <Avatar user={Teamdream.session.user} />}
        <Form error={error}>
          {Teamdream.session.isAuthenticated && <UserName user={Teamdream.session.user} />}
          <TextArea
            placeholder="Write a comment..."
            field="content"
            value={content}
            minRows={1}
            onChange={commentChanged}
            onFocus={handleOnFocus}
            inputRef={inputRef}
          />
          {content && (
            <>
              <MultiImageUploader field="attachments" maxUploads={2} previewMaxWidth={100} onChange={setAttachments} />
              <Button color="positive" onClick={submit}>
                Submit
              </Button>
            </>
          )}
        </Form>
      </div>
    </>
  );
};
