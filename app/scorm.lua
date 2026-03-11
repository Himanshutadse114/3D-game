local M = {}

-- Portable SCORM 1.2 Integration for Defold
-- This module uses html5.run to communicate with the LMS API

local function js_run(code)
    if html5 then
        return html5.run(code)
    end
    return nil
end

function M.initialize()
    print("SCORM: Initializing...")
    js_run([[
        (function() {
            var findAPI = function(win) {
                var findAttempts = 0;
                while ((win.API == null) && (win.parent != null) && (win.parent != win)) {
                    findAttempts++;
                    if (findAttempts > 500) return null;
                    win = win.parent;
                }
                return win.API;
            };

            window.scormAPI = findAPI(window);
            if (!window.scormAPI && window.opener) {
                window.scormAPI = findAPI(window.opener);
            }

            if (window.scormAPI) {
                var result = window.scormAPI.LMSInitialize("");
                if (result === "true") {
                    window.scormAPI.LMSSetValue("cmi.core.lesson_status", "incomplete");
                    window.scormAPI.LMSCommit("");
                    console.log("SCORM 1.2: Initialized successfully");
                } else {
                    console.error("SCORM 1.2: LMSInitialize failed");
                }
            } else {
                console.warn("SCORM 1.2: API not found. Running in standalone mode.");
            }
        })();
    ]])
end

function M.set_score(raw, max_score, min_score)
    raw = raw or 0
    max_score = max_score or 100
    min_score = min_score or 0
    
    js_run(string.format([[
        if (window.scormAPI) {
            window.scormAPI.LMSSetValue("cmi.core.score.raw", "%s");
            window.scormAPI.LMSSetValue("cmi.core.score.max", "%s");
            window.scormAPI.LMSSetValue("cmi.core.score.min", "%s");
            window.scormAPI.LMSCommit("");
            console.log("SCORM 1.2: Score updated to " + %s);
        }
    ]], tostring(raw), tostring(max_score), tostring(min_score), tostring(raw)))
end

function M.set_lesson_status(status)
    -- common values: "passed", "completed", "failed", "incomplete", "browsed", "not attempted"
    js_run(string.format([[
        if (window.scormAPI) {
            window.scormAPI.LMSSetValue("cmi.core.lesson_status", "%s");
            window.scormAPI.LMSCommit("");
            console.log("SCORM 1.2: Status updated to %s");
        }
    ]], status, status))
end

function M.commit()
    js_run([[
        if (window.scormAPI) {
            window.scormAPI.LMSCommit("");
        }
    ]])
end

function M.terminate()
    js_run([[
        if (window.scormAPI) {
            window.scormAPI.LMSFinish("");
            console.log("SCORM 1.2: Session terminated");
        }
    ]])
end

return M
