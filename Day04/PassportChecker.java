import java.io.BufferedReader;
import java.io.FileReader;
import java.io.IOException;
import java.util.HashMap;

public class PassportChecker {
  private static HashMap<String, Boolean> fields = new HashMap<String, Boolean>();

  public static void main(String[] args) {
    populateFields();

    BufferedReader reader;
    try {
      reader = new BufferedReader(new FileReader("./input.txt"));

      int valid = 0;
      String passport = "";
      String line = reader.readLine();
      while (line != null) {
        if (line.isEmpty()) {
          if (processPassport(passport)) {
            valid++;
          }
          passport = "";
        } else {
          if (passport.isEmpty()) {
            passport = line;
          } else {
            passport += " " + line;
          }
        }
        line = reader.readLine();
      }
      reader.close();
      if (processPassport(passport)) {
        valid++;
      }
      System.out.println("Found " + valid + " passports!");
    } catch (IOException ex) {
      ex.printStackTrace();
    }
  }

  private static void populateFields() {
    fields.put("byr", true);
    fields.put("iyr", true);
    fields.put("eyr", true);
    fields.put("hgt", true);
    fields.put("hcl", true);
    fields.put("ecl", true);
    fields.put("pid", true);
    fields.put("cid", false);
  }

  private static Boolean processPassport(String passportData) {
    HashMap<String, String> passport = new HashMap<String, String>();

    String[] parts = passportData.split(" ");
    for (String p : parts) {
      String[] kv = p.split(":");
      passport.put(kv[0], kv[1]);
    }

    for (String k : fields.keySet()) {
      if (fields.get(k)) {
        if (passport.get(k) == null) {
          return false;
        }
        switch (k) {
          case "byr":
            if (passport.get(k).length() != 4) {
              return false;
            }
            try {
              int byr = Integer.parseInt(passport.get(k));
              if (byr < 1920 || byr > 2002) {
                return false;
              }
            } catch (Exception e) {
              return false;
            }
            break;
          case "iyr":
            if (passport.get(k).length() != 4) {
              return false;
            }
            try {
              int iyr = Integer.parseInt(passport.get(k));
              if (iyr < 2010 || iyr > 2020) {
                return false;
              }
            } catch (Exception e) {
              return false;
            }
            break;
          case "eyr":
            if (passport.get(k).length() != 4) {
              return false;
            }
            try {
              int eyr = Integer.parseInt(passport.get(k));
              if (eyr < 2020 || eyr > 2030) {
                return false;
              }
            } catch (Exception e) {
              return false;
            }
            break;
          case "hgt":
            String hgt = passport.get(k);
            if (hgt.endsWith("cm")) {
              hgt = hgt.replace("cm", "");
              try {
                int h = Integer.parseInt(hgt);
                if (h < 150 || h > 193) {
                  return false;
                }
              } catch (Exception e) {
                return false;
              }
              continue;
            }
            if (hgt.endsWith("in")) {
              hgt = hgt.replace("in", "");
              try {
                int h = Integer.parseInt(hgt);
                if (h < 59 || h > 76) {
                  return false;
                }
              } catch (Exception e) {
                return false;
              }
              continue;
            }
            return false;
          case "hcl":
            if (passport.get(k).length() != 7) {
              return false;
            }
            char first = passport.get(k).charAt(0);
            if (first != '#') {
              return false;
            }
            break;
          case "ecl":
            if (passport.get(k).length() != 3) {
              return false;
            }
            String ec = passport.get(k);
            if (!ec.equals("amb") && !ec.equals("blu") && !ec.equals("brn") && !ec.equals("gry") && !ec.equals("grn")
                && !ec.equals("hzl") && !ec.equals("oth")) {
              return false;
            }
            break;
          case "pid":
            if (passport.get(k).length() != 9) {
              return false;
            }
            break;
          default:
            continue;
        }
      }
    }

    return true;
  }
}