import { MicrophoneIcon } from "@heroicons/react/24/outline";

export default function TakeNoteButton() {
  return (
    <button className="flex mr-3 px-4 py-1 text-base text-teal-600 font-semibold rounded-full border border-teal-300 hover:text-white hover:bg-teal-600 hover:border-transparent focus:outline-none">
      <MicrophoneIcon className="size-5 mr-1" />
      <span>Take a note</span>
    </button>
  );
}
