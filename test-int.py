import asyncio
import httpx # Changed from aiohttp
import json

# The endpoint where you want to send the POST requests
TARGET_ENDPOINT = "http://localhost:8380/v1/scrape" # Change this if your endpoint is different

URLS_TO_SUBMIT = [
    "https://promptrefiner.me",
    "https://propcopy.ai",
    "https://reddit.com",
    "https://wwww.delta.com",
    "https://nasdaq.com",
    "https://intoli.com/blog/not-possible-to-block-chrome-headless/chrome-headless-test.html"
]

async def post_url(client: httpx.AsyncClient, url_to_submit: str, target_endpoint: str):
    """
    Sends a POST request with the given URL in the JSON payload.
    """
    payload = {"url": url_to_submit}
    headers = {"Content-Type": "application/json"}

    print(f"Submitting task for URL: {url_to_submit} to {target_endpoint}")
    try:
        response = await client.post(target_endpoint, data=json.dumps(payload), headers=headers)
        response_text = response.text # httpx.Response.text is a property
        if response.status_code >= 200 and response.status_code < 300:
            print(f"Successfully submitted {url_to_submit}. Status: {response.status_code}. Response: {response_text[:200]}...")
        else:
            print(f"Failed to submit {url_to_submit}. Status: {response.status_code}. Response: {response_text}")
        return response.status_code, response_text
    except httpx.ConnectError as e: # Changed from aiohttp.ClientConnectorError
        print(f"Connection error for {url_to_submit} to {target_endpoint}: {e}")
        return None, str(e)
    except Exception as e:
        print(f"An unexpected error occurred while submitting {url_to_submit}: {e}")
        return None, str(e)

async def main():
    """
    Main function to create tasks and run them concurrently.
    """
    async with httpx.AsyncClient() as client: # Changed from aiohttp.ClientSession
        tasks = []
        for url in URLS_TO_SUBMIT:
            tasks.append(post_url(client, url, TARGET_ENDPOINT))
        
        results = await asyncio.gather(*tasks, return_exceptions=True) # return_exceptions=True to handle individual task failures
        
        print("\n--- Summary ---")
        for i, result in enumerate(results):
            url = URLS_TO_SUBMIT[i]
            if isinstance(result, Exception):
                print(f"Task for {url} failed with an exception: {result}")
            elif result[0] is None: # Custom handling for connection errors etc.
                 print(f"Task for {url} failed: {result[1]}")
            elif result[0] >= 200 and result[0] < 300:
                print(f"Task for {url}: Succeeded (Status {result[0]})")
            else:
                print(f"Task for {url}: Failed (Status {result[0]})")


if __name__ == "__main__":
    print(f"Starting submission of URLs to {TARGET_ENDPOINT}...")
    asyncio.run(main())
    print("All submissions attempted.")
