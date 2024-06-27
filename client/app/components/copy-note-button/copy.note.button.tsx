import { DocumentDuplicateIcon } from "@heroicons/react/24/outline";

export default function CopyNoteButton() {
  return (
    <button data-testid className="flex px-4 py-1 text-base text-teal-600 font-semibold rounded-full border border-teal-300 hover:text-white hover:bg-teal-600 hover:border-transparent focus:outline-none">
      <DocumentDuplicateIcon className="size-5 mr-1" />
      <span>Copy note</span>
    </button>
  );
}
