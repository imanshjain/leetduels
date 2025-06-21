// components/CodeEditor.jsx
'use client'

import React, { useState, ChangeEvent } from 'react';
import Editor from '@monaco-editor/react';

const languages = ['javascript', 'python', 'java', 'c', 'cpp', 'csharp', 'sql'];

const CodeEditor: React.FC = () => {
  const [language, setLanguage] = useState<string>('python');
  const [code, setCode] = useState<string>('# Write your code here');

  const handleLanguageChange = (e: ChangeEvent<HTMLSelectElement>) => {
    const selectedLang = e.target.value;
    setLanguage(selectedLang);
    setCode('');
  };

  return (
    <div className="p-4">
      <div className="m-2">
        <label htmlFor="language">Select Language: </label>
        <select
          className="bg-gray-900"
          id="language"
          value={language}
          onChange={handleLanguageChange}
        >
          {languages.map((lang) => (
            <option key={lang} value={lang}>
              {lang}
            </option>
          ))}
        </select>
      </div>

      <Editor
        height="500px"
        language={language}
        value={code}
        onChange={(value) => setCode(value || '')}
        theme="vs-dark"
        options={{
          fontSize: 14,
          minimap: { enabled: false },
        }}
      />
    </div>
  );
};

export default CodeEditor;