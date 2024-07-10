"use client";

import {
  CheckCircleIcon,
  DocumentDuplicateIcon,
} from "@heroicons/react/24/outline";
import { useEffect, useState } from "react";
import BaseButton from "../base/button/base.button";

export default function CopyNoteButton({ noteText }: { noteText: string }) {
  const [isCopied, setIsCopied] = useState(false);

  useEffect(() => {
    if (isCopied) {
      const timeout = setTimeout(() => {
        setIsCopied(false);
      }, 2000);

      return () => clearTimeout(timeout);
    }
  }, [isCopied]);

  const handleNoteCopy = () => {
    navigator.clipboard.writeText(noteText);
    setIsCopied(true);
  };

  const icon = isCopied ? (
    <CheckCircleIcon className="size-6 mr-1" />
  ) : (
    <DocumentDuplicateIcon className="size-5 mr-1" />
  );

  return (
    <BaseButton
      text={isCopied ? "Copied!" : "Copy note"}
      icon={icon}
      disabled={!noteText}
      onClick={handleNoteCopy}
      className="text-teal-600 border-teal-300 hover:bg-teal-600"
    />
  );
}
